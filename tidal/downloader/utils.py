from os import path
from posixpath import basename
from urllib.parse import unquote, urlsplit
from tidalapi.media import Track
from tidalapi.artist import Artist, Role

def url_to_filename(url: str) -> str:
    """Convert a URL to a valid filename.

    Args:
        url (str): The URL to convert.

    Returns:
        str: The corresponding filename.

    Raises:
        ValueError: If the URL contains invalid characters for a filename.
    """
    urlpath: str = urlsplit(url).path
    base_name: str = basename(unquote(urlpath))

    if path.basename(base_name) != base_name or unquote(basename(urlpath)) != base_name:
        raise ValueError  # reject '%2f' or 'dir%5Cbasename.ext' on Windows

    return base_name

def name_builder_title(media: Track) -> str:
    result: str = (
        media.full_name if media.full_name else media.name if media.name else "Unknown Title"
    )

    return result

def name_builder_artist(media: Track, delimiter: str = ", ") -> str:
    """Builds a string of artist names for a track, video, or album.

    Returns a delimited string of all artist names associated with the given media.

    Args:
        media (Track): The media object to extract artist names from.
        delimiter (str, optional): The delimiter to use between artist names. Defaults to ", ".

    Returns:
        str: A delimited string of artist names.
    """
    return delimiter.join((artist.name if artist.name else "Unknown Artist" for artist in media.artists) if media.artists else ["Unknown Artist"])


def name_builder_album_artist(media: Track, first_only: bool = False, delimiter: str = ", ") -> str:
    """Builds a string of main album artist names for a track or album.

    Returns a delimited string of main artist names from the album, optionally including only the first main artist.

    Args:
        media (Track | Album): The media object to extract artist names from.
        first_only (bool, optional): If True, only the first main artist is included. Defaults to False.
        delimiter (str, optional): The delimiter to use between artist names. Defaults to ", ".

    Returns:
        str: A delimited string of main album artist names.
    """
    artists_tmp: list[str] = []
    artists: list[Artist] = (media.album.artists if isinstance(media, Track) else media.artists) if media.album and media.album.artists else []

    for artist in artists:
        if not artist.roles:
            continue

        if Role.main in artist.roles:
            artists_tmp.append(artist.name if artist.name else "Unknown Artist")

            if first_only:
                break

    return delimiter.join(artists_tmp)