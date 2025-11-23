from tidalapi.media import Track, AudioExtensions, Stream
from tidalapi.session import Session

from pathlib import Path
from shutil import move
from datetime import datetime
from requests import Session as RequestSession, adapters, Response, get
from os import path
from ffmpeg import FFmpeg
from concurrent import futures

from downloader.metadata import Metadata
from downloader.config import DownloaderConfig
from downloader import utils

class TidalDownloader:
    config: DownloaderConfig
    session: Session
    adapter: adapters.HTTPAdapter

    def __init__(self, config: DownloaderConfig | None = None) -> None:
        """
        Initializes the downloader with the provided configuration.

        Args:
            config (DownloaderConfig): The configuration object containing settings for the downloader.
        """

        if not config:
            return

        self.config = config
        self.session = Session()
        self.session.load_oauth_session(config.auth_token_type, config.auth_access_token, config.auth_refresh_token, datetime.fromtimestamp(config.auth_expiry_time))
        self.session.audio_quality = config.audio_quality
        self.session.config.client_id = config.auth_client_id
        self.session.config.client_secret = config.auth_client_secret

        self.adapter = adapters.HTTPAdapter(max_retries=self.config.download_retries)

        if not path.exists(self.config.download_path):
            Path(self.config.download_path).mkdir(parents=True, exist_ok=True)
        
        if not path.exists(self.config.temp_path):
            Path(self.config.temp_path).mkdir(parents=True, exist_ok=True)

    def search_by_isrcs(self, isrcs: list[str]) -> list[Track]:
        """
        Search for tracks based on a list of ISRC codes.

        Args:
            isrcs (list[str]): A list of ISRC codes to search for.

        Returns:
            list[Track]: A list of Track objects matching the provided ISRC codes.
        """
        tracks: list[Track] = []

        for isrc in isrcs:
            try:
                tracks.extend(self.session.get_tracks_by_isrc(isrc))
            except:
                print(f"Could not find track with ISRC: {isrc}")
                pass

        return tracks

    def search(self, query: str, limit: int = 50, offset: int = 0) -> list[Track]:
        """
        Search for tracks based on a query string.

        Args:
            query (str): The search query string.
            limit (int, optional): The maximum number of tracks to return. Defaults to 50.
            offset (int, optional): The starting position for the search results. Defaults to 0.

        Returns:
            list[Track]: A list of Track objects matching the search query.
        """
        res = self.session.search(query, models=[Track], limit=limit, offset=offset)
        return res['tracks']

    def download(self, track: Track) -> Path | None:
        """
        Downloads the specified track.

        Args:
            track (Track): The track to be downloaded.

        Returns:
            Path | None: The path to the downloaded file, or None if the download failed.
        """
        stream = track.get_stream()
        stream_manifest = stream.get_stream_manifest()
        urls: list[str] = list(stream_manifest.get_urls())

        paths: list[Path] = []

        with futures.ThreadPoolExecutor(max_workers=self.config.download_threads) as executor:
            l_fututres: list[futures.Future] = [
                executor.submit(self._get_url, url)
                for url in urls
            ]

            for future in futures.as_completed(l_fututres):
                file_path = future.result()
                paths.append(file_path)

        temp_file = Path(self.config.temp_path) / f"{track.artist.name if track.artist else 'Unknown Artist'} - {track.name}"

        self._combine_files(paths, temp_file)

        [path.unlink(missing_ok=True) for path in paths]

        flac_file = self._extract_flac(temp_file)

        temp_file.unlink(missing_ok=True)

        self._handle_metadata(track, stream, flac_file)

        final_file = self._move_file(flac_file, track)

        return final_file

    def _move_file(self, file_path: Path, track: Track) -> Path:
        assert track.artist is not None
        assert track.artist.name is not None
        assert track.album is not None
        assert track.album.name is not None

        new_path = Path(self.config.download_path) / track.artist.name / track.album.name / f"{track.track_num:02d} - {track.name}{AudioExtensions.FLAC}"
        
        new_path.parent.mkdir(parents=True, exist_ok=True)

        move(file_path, new_path)

        return new_path

    def _handle_metadata(self, track: Track, stream: Stream, path_media: Path) -> None:
        """
        Handles the metadata for a given track and stream, and saves it to the specified media path.

        Args:
            track (Track): The track object containing metadata such as title, album, artists, etc.
            stream (Stream): The stream object containing audio-related metadata such as replay gain and peak amplitude.
            path_media (Path): The file path where the media file is located.

        Returns:
            None

        Metadata Processed:
            - Title, album, artists, album artist, and track number.
            - Release date, copyright, and ISRC.
            - Lyrics (both synced and unsynced).
            - Cover image data.
            - Replay gain and peak amplitude for both album and track.
            - Explicit content flag.
            - UPC and share URL.

        Notes:
            - Attempts to retrieve lyrics from the track object. If unavailable, logs a message.
            - Retrieves album cover image data from the provided URL.
            - Saves the metadata using the `Metadata` class.
        """
        if not track.album:
            return

        release_date: str = (
            track.album.available_release_date.strftime("%Y-%m-%d")
            if track.album.available_release_date
            else track.album.release_date.strftime("%Y-%m-%d") if track.album.release_date else ""
        )
        copy_right: str = track.copyright if hasattr(track, "copyright") and track.copyright else ""
        isrc: str = track.isrc if hasattr(track, "isrc") and track.isrc else ""
        explicit: bool = track.explicit if hasattr(track, "explicit") else False
        title = utils.name_builder_title(track)
        lyrics_synced: str = ""
        lyrics_unsynced: str = ""
        cover_data: bytes | None = None

        # Try to retrieve lyrics.
        try:
            lyrics_obj = track.lyrics()

            if lyrics_obj.text:
                lyrics_unsynced = lyrics_obj.text
            if lyrics_obj.subtitles:
                lyrics_synced = lyrics_obj.subtitles
        except:
            print(f"Could not retrieve lyrics.")

        url_cover = track.album.image(1280)
        cover_data = self._cover_data(url=url_cover)

        m: Metadata = Metadata(
            path_file=path_media,
            target_upc={"MP3": "UPC", "MP4": "UPC", "FLAC": "UPC"},
            lyrics=lyrics_synced,
            lyrics_unsynced=lyrics_unsynced,
            copy_right=copy_right,
            title=title,
            artists=utils.name_builder_artist(track, delimiter=self.config.metadata_delimiter_artist),
            album=track.album.name if track.album and track.album.name else "",
            tracknumber=track.track_num,
            date=release_date,
            isrc=isrc,
            albumartist=utils.name_builder_album_artist(track, delimiter=self.config.metadata_delimiter_album_artist),
            totaltrack=track.album.num_tracks if track.album and track.album.num_tracks else 1,
            totaldisc=track.album.num_volumes if track.album and track.album.num_volumes else 1,
            discnumber=track.volume_num if track.volume_num else 1,
            cover_data=cover_data if cover_data else b"",
            album_replay_gain=stream.album_replay_gain,
            album_peak_amplitude=stream.album_peak_amplitude,
            track_replay_gain=stream.track_replay_gain,
            track_peak_amplitude=stream.track_peak_amplitude,
            url_share=track.share_url,
            replay_gain_write=True,
            upc=track.album.upc if track.album and track.album.upc else "",
            explicit=explicit,
        )

        m.save()

    def _cover_data(self, url: str | None = None, path_file: str | None = None) -> bytes | None:
        """Retrieve cover image data from a URL or file.

        Args:
            url (str | None, optional): URL to download image from. Defaults to None.
            path_file (str | None, optional): Path to image file. Defaults to None.

        Returns:
            bytes | None: Image data or None on failure.
        """
        result: bytes | None = None

        if url:
            response: Response | None = None
            try:
                response = get(url, timeout=self.config.download_timeout)
                result = response.content
            except Exception as e:
                print(e)
            finally:
                if response:
                    response.close()
        elif path_file:
            try:
                with open(path_file, "rb") as f:
                    result = f.read()
            except OSError as e:
                print(e)

        return result

    def _extract_flac(self, file_path: Path) -> Path:
        """
        Extracts a FLAC audio file from the given input file using FFmpeg.
        This method takes an input audio file, processes it using FFmpeg, and
        outputs a FLAC file with the same base name as the input file.
        Args:
            file_path (Path): The path to the input audio file.
        Returns:
            Path: The path to the extracted FLAC file.
        Raises:
            FFmpegError: If the FFmpeg execution fails.
        """
        path_media_out = file_path.with_suffix(AudioExtensions.FLAC)
    
        ffmpeg = (
            FFmpeg(executable=self.config.ffmpeg_binary_path)
            .input(url=file_path)
            .output(
                url=path_media_out,
                map=0,
                movflags="use_metadata_tags",
                acodec="copy",
                map_metadata="0:g",
                loglevel="quiet",
            )
        )

        ffmpeg.execute()

        return path_media_out

    def _get_url(self, url: str) -> Path:
        """
        Downloads a file from the given URL and saves it to a temporary path.

        Args:
            url (str): The URL of the file to download.

        Returns:
            Path: The path to the downloaded file.

        Raises:
            requests.exceptions.RequestException: If there is an issue with the HTTP request.
            HTTPError: If the HTTP request returns an unsuccessful status code.

        Notes:
            - The file is downloaded in chunks, with the chunk size determined by `self.config.block_size`.
            - The temporary path for the file is determined by `self.config.temp_path`.
            - The download timeout is specified by `self.config.download_timeout`.
        """
        file_path = Path(self.config.temp_path) / utils.url_to_filename(url)

        with RequestSession() as session:
            session.mount("https://", self.adapter)
            res = session.get(url, stream=True, timeout=self.config.download_timeout)
            res.raise_for_status()

            with open(file_path, "wb") as f:
                for chunk in res.iter_content(chunk_size=self.config.block_size):
                    if not chunk:
                        break

                    f.write(chunk)

        return file_path

    def _combine_files(self, file_parts: list[Path], output_file: Path) -> None:
        """
        Combines multiple file parts into a single output file.

        Args:
            file_parts (list[Path]): A list of Path objects representing the file parts to be combined.
            output_file (Path): The Path object representing the output file where the combined content will be written.

        Returns:
            None
        """
        file_parts.sort(key=lambda p: int(str(p.stem).split("_")[-1]))

        with open(output_file, "wb") as outfile:
            for part in file_parts:
                with open(part, "rb") as infile:
                    while segment := infile.read(self.config.block_size):
                        outfile.write(segment)