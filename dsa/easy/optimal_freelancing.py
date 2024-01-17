def optimalFreelancing(jobs: list[dict[str, int]]):
    days = 7
    profit = 0
    jobs.sort(key=lambda x: x["payment"], reverse=True)
    timeline = [0] * days
    for job in jobs:
        max_time = min((job["deadline"], days))
        for time in reversed(range(max_time)):
            if not timeline[time]:
                timeline[time] = True
                profit += job["payment"]
                break
    return profit


out = optimalFreelancing(
    [
        {"deadline": 1, "payment": 1},
        {"deadline": 2, "payment": 1},
        {"deadline": 2, "payment": 2},
    ]
)

print(out)
assert out == 3
