# Do not edit the class below except for
# the constructor and the createSet, find,
# and union methods. Feel free to add new
# properties and methods to the class.
class UnionFind:
    def __init__(self):
        self.tree = {}
        self.ranks = {}

    def createSet(self, value) -> None:
        self.tree[value] = value
        self.ranks[value] = 0

    def find(self, value):
        if value not in self.tree:
            return None
        if value != self.tree[value]:
            self.tree[value] = self.find(self.tree[value])
        return self.tree[value]

    def union(self, val_1, val_2) -> None:
        if val_1 not in self.tree or val_2 not in self.tree:
            return
        val_1_root = self.find(val_1)
        val_2_root = self.find(val_2)
        if self.ranks[val_1_root] < self.ranks[val_2_root]:
            self.tree[val_1_root] = val_2_root
        elif self.ranks[val_1_root] > self.ranks[val_2_root]:
            self.tree[val_2_root] = val_1_root
        else:
            self.tree[val_2_root] = val_1_root
            self.ranks[val_1_root] += 1
