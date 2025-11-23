from models.track import Track
from downloader.tidal import TidalDownloader
from pathlib import Path
from .compare import Comparator
from tidalapi.media import Quality
from requests import get
from utils import get_env_var

from models.config import GeneralConfig, TidalConfig

from downloader.tidal import TidalDownloader
from downloader.config import DownloaderConfig

class Downloader:
    enabled: bool
    tidal: TidalDownloader
    comparator: Comparator

    def __init__(self, enabled: bool, tidal: TidalDownloader, comparator: Comparator):
        self.enabled = enabled
        self.tidal = tidal
        self.comparator = comparator

    def download(self, track: Track) -> Path | None:
        if not self.enabled:
            return None

        print("Downloading ", track)

        if len(track.ISRCs) > 0:
            tracks = self.tidal.search_by_isrcs(track.ISRCs)
        else:
            tracks = self.tidal.search(f"{track.Artist} {track.Title}")

        closest = self.comparator.get_closest_track(track, tracks)

        if closest:
            return self.tidal.download(closest)

def NewDownloader(comparator: Comparator) -> Downloader:
    res = get("{}/config/tidal".format(get_env_var("CONFIG_SERVICE_URL")))
    if res.status_code != 200:
        raise ValueError("Failed to load Tidal config")

    body = res.json()

    assert "General" in body
    assert "Tidal" in body

    general_cfg = GeneralConfig(body["General"])
    tidal_cfg = TidalConfig(body["Tidal"])

    if tidal_cfg.Enabled:
        cfg = DownloaderConfig(
            auth_token_type=tidal_cfg.AuthTokenType,
            auth_access_token=tidal_cfg.AuthAccessToken,
            auth_refresh_token=tidal_cfg.AuthRefreshToken,
            auth_expiry_time=tidal_cfg.AuthExpiresAt,
            auth_client_id=tidal_cfg.AuthClientID,
            auth_client_secret=tidal_cfg.AuthClientSecret,
            download_timeout=tidal_cfg.DownloadTimeout,
            download_retries=tidal_cfg.DownloadRetries,
            download_threads=tidal_cfg.DownloadThreads,
            audio_quality=Quality(tidal_cfg.AudioQuality),
            file_permissions=tidal_cfg.FilePermissions,
            directory_permissions=tidal_cfg.DirectoryPermissions,
            owner=tidal_cfg.Owner,
            group=tidal_cfg.Group,
            download_path=general_cfg.DownloadPath,
            temp_path=general_cfg.TempPath,
        )

        tidal = TidalDownloader(cfg)
    else:
        print("Tidal downloader is disabled")

        tidal = TidalDownloader()

    return Downloader(
        enabled=tidal_cfg.Enabled,
        tidal=tidal,
        comparator=comparator
    )
