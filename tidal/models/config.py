from datetime import datetime

class GeneralConfig:
    DownloadPath: str
    TempPath: str

    def __init__(self, data: dict) -> None:
        assert "DownloadPath" in data
        assert "TempPath" in data

        self.DownloadPath = data["DownloadPath"]
        self.TempPath = data["TempPath"]

    def __str__(self) -> str:
        return f"GeneralConfig(DownloadPath={self.DownloadPath}, TempPath={self.TempPath})"

class TidalConfig:
    Enabled: bool
    AuthTokenType: str
    AuthAccessToken: str
    AuthRefreshToken: str
    AuthExpiresAt: float
    AuthClientID: str
    AuthClientSecret: str
    DownloadTimeout: int
    DownloadRetries: int
    DownloadThreads: int
    AudioQuality: str
    FilePermissions: int
    DirectoryPermissions: int
    Owner: int
    Group: int

    def __init__(self, data: dict) -> None:
        assert "Enabled" in data
        assert "AuthTokenType" in data
        assert "AuthAccessToken" in data
        assert "AuthRefreshToken" in data
        assert "AuthExpiresAt" in data
        assert "AuthClientID" in data
        assert "AuthClientSecret" in data
        assert "DownloadTimeout" in data
        assert "DownloadRetries" in data
        assert "DownloadThreads" in data
        assert "AudioQuality" in data
        assert "FilePermissions" in data
        assert "DirectoryPermissions" in data
        assert "Owner" in data
        assert "Group" in data

        self.Enabled = data["Enabled"]
        self.AuthTokenType = data["AuthTokenType"]
        self.AuthAccessToken = data["AuthAccessToken"]
        self.AuthRefreshToken = data["AuthRefreshToken"]
        self.AuthExpiresAt = datetime.strptime(data["AuthExpiresAt"], "%Y-%m-%dT%H:%M:%SZ").timestamp()
        self.AuthClientID = data["AuthClientID"]
        self.AuthClientSecret = data["AuthClientSecret"]
        self.DownloadTimeout = data["DownloadTimeout"]
        self.DownloadRetries = data["DownloadRetries"]
        self.DownloadThreads = data["DownloadThreads"]
        self.AudioQuality = data["AudioQuality"]
        self.FilePermissions = data["FilePermissions"]
        self.DirectoryPermissions = data["DirectoryPermissions"]
        self.Owner = data["Owner"]
        self.Group = data["Group"]

    def __str__(self) -> str:
        return f"TidalConfig(AuthTokenType={self.AuthTokenType}, AuthAccessToken={self.AuthAccessToken}, AuthRefreshToken={self.AuthRefreshToken}, AuthExpiresAt={self.AuthExpiresAt}, AuthClientID={self.AuthClientID}, AuthClientSecret={self.AuthClientSecret}, DownloadTimeout={self.DownloadTimeout}, DownloadRetries={self.DownloadRetries}, DownloadThreads={self.DownloadThreads}, AudioQuality={self.AudioQuality}, FilePermissions={self.FilePermissions}, DirectoryPermissions={self.DirectoryPermissions}, Owner={self.Owner}, Group={self.Group})"