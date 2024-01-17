from typing import Dict, List

BASE_TYPE = List[List[int]]


def stableInternships(interns: BASE_TYPE, teams: BASE_TYPE) -> BASE_TYPE:
    """
    using the gale-shapley algorithm
    """
    unmatched = list(range(len(interns)))
    team_prefs = [{x: i for i, x in enumerate(team)} for team in teams]
    intern_prefs = [0] * len(interns)
    pairs: Dict[int, int] = {}

    while unmatched:
        curr = unmatched.pop()
        team = interns[curr][intern_prefs[curr]]
        intern_prefs[curr] += 1

        if team not in pairs:
            pairs[team] = curr
            continue
        if team_prefs[team][curr] < team_prefs[team][pairs[team]]:
            unmatched.append(pairs[team])
            pairs[team] = curr
            continue
        unmatched.append(curr)

    return [[intern, team] for team, intern in pairs.items()]
