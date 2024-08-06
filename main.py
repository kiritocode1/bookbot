
def get_letter_count(words):
    letter_count = Counter()
    for word in words:
        for letter in word.lower():
            letter_count[letter] += 1
    return letter_count


from collections import Counter
with open("books/frankenstein.txt", "r") as f:
    file_contents = f.read()
    words = list(file_contents.split())
    count =list(get_letter_count(words).items())
    # for letter, counts in count:
    #     print(f"the letter '{letter}' character was found {counts} times")
    count.sort(key=lambda x: x[1], reverse=True)
    for letter, counts in count:
        print(f"the letter '{letter}' character was found {counts} times")