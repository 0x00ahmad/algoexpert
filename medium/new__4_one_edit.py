def one_edit_remove(s1, s2):
    i, j = 0, 0
    while i < len(s1) and j < len(s2):
        if s1[i] == s2[j]:
            i, j = i + 1, j + 1
            continue
        if i != j:
            return False
        i += 1
    return True


def one_edit_add(s1, s2):
    i, j = 0, 0
    while i < len(s1) and j < len(s2):
        if s1[i] == s2[j]:
            i, j = i + 1, j + 1
            continue
        if i != j:
            return False
        j += 1
    return True


def one_edit_replace(s1, s2):
    found_difference = False
    for i in range(len(s1)):
        if s1[i] != s2[i]:
            if found_difference:
                return False
            found_difference = True
    return True


def oneEdit(s1: str, s2: str) -> bool:
    """
    Given two strings, determine if they are one edit away from each other.
    """
    if s1 == s2:
        return True
    if len(s1) == len(s2):
        return one_edit_replace(s1, s2)
    if len(s1) == len(s2) + 1:
        return one_edit_remove(s1, s2)
    if len(s1) == len(s2) - 1:
        return one_edit_add(s1, s2)
    return False
