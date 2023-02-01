def reverse_words(sentence):
    # write you code
    # nggak bisa simply pakai split, karena character "  " juga bagian dari kalimat

    # words = sentence.split('\s')

    def split_sentences(sentence):
        reversed = []
        # apa beda [] dan list ?
        j = 0
        temp = []

        for i in range (len(sentence)-1, -1, -1):
            if sentence[i] != " ":
                temp.append(sentence[i])
            else:
                temp = []

            reversed.append(temp)
            j += 1
        print(f"reversed {reversed}")
        return "ok ok"

    words = split_sentences(sentence)
    print(words)

    # left = 0
    # right = len(words)-1
    # while left < right:
    #     words[left], words[right] = words[right], words[left]
    #     left += 1
    #     right -= 1
    # return " ".join(words)

if __name__ == "__main__":
    test = [
        "We love Python",
        "To be or not to be",
        "You are amazing",
        "Hello     World",
        "Hey"
    ]
    exp = [
        "Python love We",
        "be to not or be To",
        "amazing are You",
        "World     Hello",
        "Hey"
    ]
    res = reverse_words(test[3])
    print(f"expected {exp[3]}, result {res}")