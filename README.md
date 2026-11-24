# Grid Unique Paths

> Striver SDE Sheet 017 · Day 3 Arrays — no DP table needed: every path is a shuffle of downs and rights, so the count is one binomial coefficient. Go, stdlib only, tested.

## Problem

Count the distinct paths from the top-left to the bottom-right corner of an m×n grid when the only moves are right and down.

## Approach

Each path makes exactly m−1 down-moves and n−1 right-moves in some order; choosing *which* of the m+n−2 steps are downs fully determines the path. So the answer is C(m+n−2, m−1) — combinatorics replaces the O(m·n) DP. The coefficient is computed by the running product `res = res·(total−k+i)/i`, which stays integral at every step because each prefix is itself a binomial coefficient.

**Complexity:** O(min(m,n)) time, O(1) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
