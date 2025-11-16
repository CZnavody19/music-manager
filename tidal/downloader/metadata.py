import pathlib

from mutagen import flac, mp4
from mutagen._file import File
from mutagen.id3._specs import PictureType

class Metadata:
    path_file: str | pathlib.Path
    title: str
    album: str
    albumartist: str
    artists: str
    copy_right: str
    tracknumber: int
    discnumber: int
    totaldisc: int
    totaltrack: int
    date: str
    composer: str
    isrc: str
    lyrics: str
    lyrics_unsynced: str
    path_cover: str
    cover_data: bytes
    album_replay_gain: float
    album_peak_amplitude: float
    track_replay_gain: float
    track_peak_amplitude: float
    url_share: str
    replay_gain_write: bool
    upc: str
    target_upc: dict[str, str]
    explicit: bool
    m: flac.FLAC

    def __init__(
        self,
        path_file: str | pathlib.Path,
        target_upc: dict[str, str],
        album: str = "",
        title: str = "",
        artists: str = "",
        copy_right: str = "",
        tracknumber: int = 0,
        discnumber: int = 0,
        totaltrack: int = 0,
        totaldisc: int = 0,
        composer: str = "",
        isrc: str = "",
        albumartist: str = "",
        date: str = "",
        lyrics: str = "",
        lyrics_unsynced: str = "",
        cover_data: bytes = b"",
        album_replay_gain: float = 1.0,
        album_peak_amplitude: float = 1.0,
        track_replay_gain: float = 1.0,
        track_peak_amplitude: float = 1.0,
        url_share: str = "",
        replay_gain_write: bool = True,
        upc: str = "",
        explicit: bool = False,
    ):
        self.path_file = path_file
        self.title = title
        self.album = album
        self.albumartist = albumartist
        self.artists = artists
        self.copy_right = copy_right
        self.tracknumber = tracknumber
        self.discnumber = discnumber
        self.totaldisc = totaldisc
        self.totaltrack = totaltrack
        self.date = date
        self.composer = composer
        self.isrc = isrc
        self.lyrics = lyrics
        self.lyrics_unsynced = lyrics_unsynced
        self.cover_data = cover_data
        self.album_replay_gain = album_replay_gain
        self.album_peak_amplitude = album_peak_amplitude
        self.track_replay_gain = track_replay_gain
        self.track_peak_amplitude = track_peak_amplitude
        self.url_share = url_share
        self.replay_gain_write = replay_gain_write
        self.upc = upc
        self.target_upc = target_upc
        self.explicit = explicit
        self.m = File(self.path_file) # pyright: ignore[reportAttributeAccessIssue]

    def _cover(self) -> bool:
        result: bool = False

        if self.cover_data:
            if isinstance(self.m, flac.FLAC):
                flac_cover = flac.Picture()
                flac_cover.type = PictureType.COVER_FRONT
                flac_cover.data = self.cover_data
                flac_cover.mime = "image/jpeg"

                self.m.clear_pictures()
                self.m.add_picture(flac_cover)

            result = True

        return result

    def save(self):
        if not self.m.tags:
            self.m.add_tags()

        if isinstance(self.m, flac.FLAC):
            self.set_flac()

        self._cover()
        self.cleanup_tags()
        self.m.save()

        return True

    def set_flac(self):
        if type(self.m.tags) is not flac.VCFLACDict:
            return

        self.m.tags["TITLE"] = self.title
        self.m.tags["ALBUM"] = self.album
        self.m.tags["ALBUMARTIST"] = self.albumartist
        self.m.tags["ARTIST"] = self.artists
        self.m.tags["COPYRIGHT"] = self.copy_right
        self.m.tags["TRACKNUMBER"] = str(self.tracknumber)
        self.m.tags["TRACKTOTAL"] = str(self.totaltrack)
        self.m.tags["DISCNUMBER"] = str(self.discnumber)
        self.m.tags["DISCTOTAL"] = str(self.totaldisc)
        self.m.tags["DATE"] = self.date
        self.m.tags["COMPOSER"] = self.composer
        self.m.tags["ISRC"] = self.isrc
        self.m.tags["LYRICS"] = self.lyrics
        self.m.tags["UNSYNCEDLYRICS"] = self.lyrics_unsynced
        self.m.tags["URL"] = self.url_share
        self.m.tags[self.target_upc["FLAC"]] = self.upc

        if self.replay_gain_write:
            self.m.tags["REPLAYGAIN_ALBUM_GAIN"] = str(self.album_replay_gain)
            self.m.tags["REPLAYGAIN_ALBUM_PEAK"] = str(self.album_peak_amplitude)
            self.m.tags["REPLAYGAIN_TRACK_GAIN"] = str(self.track_replay_gain)
            self.m.tags["REPLAYGAIN_TRACK_PEAK"] = str(self.track_peak_amplitude)

    def cleanup_tags(self):
        if type(self.m.tags) is not flac.VCFLACDict:
            return
        
        for key, value in self.m.tags.items():
            if value == "" or value == [""]:
                del self.m.tags[key]