import re

textMatch = re.compile('[0-9]+(?:s|p|d|f)\^[0-9]+', re.MULTILINE)

with open('subkulit.txt') as f:
    text = f.read()
    text = re.findall(textMatch, text)

    print(text)
    f.close()