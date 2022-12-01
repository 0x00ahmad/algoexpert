def reverseWordsInString(string: str) -> str:
    reversed_word = ""
    current_word = ""
    for i in range(len(string) - 1, -1, -1):
        character = string[i]
        is_space = character == " "
        if is_space:
            reversed_word += current_word + " "
        current_word = "" if is_space else character + current_word
    return reversed_word + current_word


# the cheaty way with python
def reverseWordsInStringTheEasyWay(string: str) -> str:
    return " ".join(reversed(string.split(" ")))
