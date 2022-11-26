def balancedBrackets(s: str) -> bool:
    stack = []
    for bracket in s:
        if bracket not in ["(", "[", "{", ")", "]", "}"]:
            continue
        is_opening_bracket = bracket in ["(", "[", "{"]
        if len(stack) == 0:
            if is_opening_bracket:
                stack.append(bracket)
            else:
                return False
        elif is_opening_bracket:
            stack.append(bracket)
        elif (
            stack[-1] == "("
            and bracket == ")"
            or stack[-1] == "["
            and bracket == "]"
            or stack[-1] == "{"
            and bracket == "}"
        ):
            stack.pop()
        else:
            return False
    return len(stack) == 0
