# Do not edit the class below except for the
# populateSuffixTrieFrom and contains methods.
# Feel free to add new properties and methods
# to the class.
class SuffixTrie:
    def __init__(self, string):
        self.root = {}
        self.endSymbol = "*"
        self.populateSuffixTrieFrom(string)

    def populateSuffixTrieFrom(self, string) -> None:
        s_len = len(string)
        for i in range(s_len):
            self.root = self.recurse(string[i:], self.root)

    def recurse(self, substring, current_hashmap):
        if not substring:
            current_hashmap[self.endSymbol] = True
            return current_hashmap
        d = (
            dict()
            if not substring[0] in current_hashmap
            else current_hashmap[substring[0]]
        )
        current_hashmap[substring[0]] = self.recurse(substring[1:], d)
        return current_hashmap

    def contains(self, string) -> bool:
        current = self.root
        for character in string:
            if character not in current:
                return False
            current = current[character]
        return self.endSymbol in current


def main():
    trie = SuffixTrie("babc")
    print(trie.root)


if __name__ == "__main__":
    main()
