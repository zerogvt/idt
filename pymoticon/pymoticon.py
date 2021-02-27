#!/usr/bin/env python3
'''
Tokenize a string keeping emoticons and stripping punctuation.
'''
import sys
import html

EMOTICONS_DB = 'emoticons.txt'


def init_emo_db():
    '''
    Creates a key-value store of emoticons.
    Keeps the first txt explanation only.
    E.g. for (noted tabs)
    :-) Smiling, happy faces; don't take me too seriously
    /tab Your basic smiley
    /tab Ha ha
    /tab Comedy
    key = ':-)', value = 'Smiling, happy faces; don't take me too seriously'
    '''
    emo_db = {}
    with open(EMOTICONS_DB, 'r') as f:
        for ln in f.readlines():
            # skip secondary explanations
            if ln.startswith("\t"):
                continue
            tokens = ln.split()
            if len(tokens) < 2:
                continue
            # first token is emoticon
            key = html.unescape(tokens[0]).strip()
            # the rest are the explanation
            val = ' '.join(tokens[1:]).strip()
            if key:
                emo_db[key] = val
    return emo_db


def emotistrip(inp, emo_db):
    '''
    Splits it input string inp into a list of words.
    Strips all punctuation except for emoticons.
    Returns a list with tokens.
    '''
    res = []
    for token in inp.split():
        # if it is an emoticon just push it on the list
        if token in emo_db:
            res.append(token)
            continue
        # if this is a word clean it up from punctuation
        cleaned = ''.join(filter(lambda ch: ch not in "()?.!/;:#<>", token))
        if len(cleaned):
            res.append(cleaned)
    return res


if __name__ == "__main__":
    emo_db = init_emo_db()
    for inp in sys.argv[1:]:
        print(emotistrip(inp, emo_db))
