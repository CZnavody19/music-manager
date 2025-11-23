from tidalapi.media import Quality
from dataclasses import dataclass

@dataclass
class DownloaderConfig:
    auth_token_type: str = "Bearer"
    auth_access_token: str = ""
    auth_refresh_token: str = ""
    auth_expiry_time: float = 0
    auth_client_id: str = ""
    auth_client_secret: str = ""

    download_timeout: int = 30
    download_retries: int = 3
    download_path: str = "./downloads"
    temp_path: str = "./temp"
    audio_quality: Quality = Quality(Quality.hi_res_lossless)

    block_size: int = 1024 * 1024  # 1 MB
    ffmpeg_binary_path: str = "/usr/bin/ffmpeg"
    download_threads: int = 4

    metadata_delimiter_artist: str = ", "
    metadata_delimiter_album_artist: str = ", "

    file_permissions: int = 0o664
    directory_permissions: int = 0o775
    owner: int = 0
    group: int = 0