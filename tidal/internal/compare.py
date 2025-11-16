from models.track import Track
from tidalapi.media import Track as TidalTrack
from thefuzz.fuzz import token_sort_ratio

class Comparator:
    WEIGHT_TITLE = 0.4
    WEIGHT_ARTIST = 0.3
    WEIGHT_LENGTH = 0.2
    WEIGHT_QUALITY = 0.1

    QUALITY_MAP = {
        "": 0,
        "LOW": 0.1,
        "HIGH": 0.2,
        "LOSSLESS": 0.8,
        "HI_RES_LOSSLESS": 1,
    }

    def get_track_score(self, original: Track, compare: TidalTrack) -> float:
        print("title:", original.Title, "vs", compare.full_name)
        print("artist:", original.Artist, "vs", compare.artist.name if compare.artist else "")
        print("length:", original.Length/1000, "vs", compare.duration)
        print("quality:", compare.audio_quality)

        t_s = token_sort_ratio(original.Title, compare.full_name) / 100
        a_s = token_sort_ratio(original.Artist, compare.artist.name if compare.artist else "") / 100
        l_s = 1 - (abs(original.Length/1000 - compare.duration) / original.Length/1000) if compare.duration else 0
        q_s = self.QUALITY_MAP.get(compare.audio_quality if compare.audio_quality else "", 0)

        print(f"Title score: {t_s}, Artist score: {a_s}, Length score: {l_s}, Quality score: {q_s}")

        return (t_s*self.WEIGHT_TITLE) + (a_s*self.WEIGHT_ARTIST) + (l_s*self.WEIGHT_LENGTH) + (q_s*self.WEIGHT_QUALITY)

    def get_closest_track(self, original: Track, candidates: list[TidalTrack]) -> TidalTrack | None:
        best_score = 0
        best_track = None

        for candidate in candidates:
            score = self.get_track_score(original, candidate)
            if score > best_score:
                best_score = score
                best_track = candidate

        return best_track
    
def NewComparator() -> Comparator:
    return Comparator()