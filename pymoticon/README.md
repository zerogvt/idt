# Part 2 - python reg exp

Given the input string: "Hello world! How (sp?) are you today (;" split it into a list of words. You must strip all punctuation except for emoticons. Result for example: ['Hello','world', 'how', 'sp', 'are', 'you', 'today', '(;']

# Preconditions
python 3.8.0 or later

# Example runs
```
pymoticon$ chmod +x pymoticon.py

pymoticon$ ./pymoticon.py "Hello world! How (sp?) are you today (;"
['Hello', 'world', 'How', 'sp', 'are', 'you', 'today', '(;']

pymoticon$ ./pymoticon.py "don't like this sth is fishy (O--<"
["don't", 'like', 'this', 'sth', 'is', 'fishy', '(O--<']

pymoticon$ ./pymoticon.py "smile! we are (ok) :-)"
['smile', 'we', 'are', 'ok', ':-)']

pymoticon$ ./pymoticon.py "I am a gorilla [:=8{]"
['I', 'am', 'a', 'gorilla', '[:=8{]']
```

## Note: on text db of emoticons
List was taken of http://marshall.freeshell.org/smileys.html
`wget http://marshall.freeshell.org/smileys.html`
I then copy pasted the mid section of html where the emoticons are to emoticons.txt
