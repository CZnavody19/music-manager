from uuid import UUID

class Track:
    ID: UUID
    Title: str
    Artist: str
    Length: int
    ISRCs: list[str]
    LinkedYoutube: bool
    LinkedPlex: bool

    def __init__(self, data: dict) -> None:
        assert "ID" in data
        assert "Title" in data
        assert "Artist" in data
        assert "Length" in data
        assert "ISRCs" in data
        assert "LinkedYoutube" in data
        assert "LinkedPlex" in data

        self.ID = UUID(data["ID"])
        self.Title = data["Title"]
        self.Artist = data["Artist"]
        self.Length = data["Length"]
        self.ISRCs = data["ISRCs"] if data["ISRCs"] is not None else []
        self.LinkedYoutube = data["LinkedYoutube"]
        self.LinkedPlex = data["LinkedPlex"]

    def __str__(self) -> str:
        return f"Track(ID={self.ID}, Title={self.Title}, Artist={self.Artist}, Length={self.Length})"
    
    def json(self) -> dict:
        return {
            "ID": str(self.ID),
            "Title": self.Title,
            "Artist": self.Artist,
            "Length": self.Length,
            "ISRCs": self.ISRCs,
            "LinkedYoutube": self.LinkedYoutube,
            "LinkedPlex": self.LinkedPlex
        }