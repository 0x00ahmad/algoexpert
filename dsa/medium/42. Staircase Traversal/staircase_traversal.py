def staircaseTraversal(height: int, maxSteps: int) -> int:
    ways = 0
    possible_ways_to_top = [1]
    for curr_height in range(height):
        window_start = curr_height - maxSteps
        window_end = curr_height
        if window_start >= 0:
            ways -= possible_ways_to_top[window_start]
        ways += possible_ways_to_top[window_end]
        possible_ways_to_top.append(ways)
    return possible_ways_to_top[height]
