import sys
emo_db = {}


def emoticons_db():
    '''creates a 'db' of emoticons'''
    with open('emoticons.txt', 'r') as f:
        emoticons = f.readlines()
        for em in emoticons:
            emo_db[em.strip()]


def emotistrip(inp):
    '''
    splits it input string inp into a list of words.
    Strips all punctuation except for emoticons
    '''
    res = []
    for token in inp.split():
        # if it is an emoticon just push it on the list
        if token in emo_db:
            res.append(token)
            continue
        # if this is a word clean it up from punctuation
        cleaned = filter(lambda ch: ch not in "()?.!/;:#<>", token)
        if len(cleaned):
            res.append(cleaned)
    return res


if __name__ == "__main__":
    for inp in sys.argv[1:]:
        print(emotistrip(inp))
