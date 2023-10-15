import yake
from typing import List


def extraction(text: str) -> List[str]:
    kw_extractor = yake.KeywordExtractor()
    language = "ru"
    max_ngram_size = 3
    numOfKeywords = 5
    custom_kw_extractor = yake.KeywordExtractor(lan=language, n=max_ngram_size, top=numOfKeywords, features=None)
    keywords = custom_kw_extractor.extract_keywords(text)
    ans = []
    for kw, score in keywords:
        ans.append(kw)
    return ans