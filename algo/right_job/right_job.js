const assert = require('assert')

const movieSet = [
    {
        name: "Movie 1",
        start: 3,
        end: 8
    }, {
        name: "Movie 2",
        start: 10,
        end: 16
    },
    {
        name: "Movie 3",
        start: 1,
        end: 6
    }, {
        name: "Movie 4",
        start: 7,
        end: 11
    }, {
        name: "Movie 5",
        start: 13,
        end: 17
    }, {
        name: "Movie 6",
        start: 2,
        end: 4
    },
    {
        name: "Movie 7",
        start: 5,
        end: 9
    }, {
        name: "Movie 8",
        start: 12,
        end: 14
    }, {
        name: "Movie 9",
        start: 15,
        end: 18
    },
]

function moviesOverlap(movie1, movie2) {
    if (movie1.length != undefined) {
        for (const movie of movie1) {
            const result = moviesOverlap(movie, movie2)
            if (result) {
                return true
            }
        }
    }
    return (movie1.start <= movie2.start && movie1.end > movie2.start) ||
        (movie1.start > movie2.start && movie1.start < movie2.end)
}

assert.equal(moviesOverlap({ start: 10, end: 15 }, { start: 10, end: 12 }), true)
assert.equal(moviesOverlap({ start: 9, end: 15 }, { start: 10, end: 12 }), true)
assert.equal(moviesOverlap({ start: 11, end: 15 }, { start: 10, end: 12 }), true)
assert.equal(moviesOverlap({ start: 11, end: 15 }, { start: 9, end: 18 }), true)
assert.equal(moviesOverlap({ start: 11, end: 15 }, { start: 15, end: 18 }), false)
assert.equal(moviesOverlap({ start: 11, end: 15 }, { start: 15, end: 18 }), false)
assert.equal(moviesOverlap([{ start: 11, end: 15 }, { start: 10, end: 12 }], { start: 10, end: 11 }), true)

function findPossibleJobsSet(remainingMovieSet, accumulatedMovieSet = []) {
    if (remainingMovieSet.length === 0) {
        return [accumulatedMovieSet]
    }
    const [head, ...tail] = remainingMovieSet
    let result1 = []
    if (!moviesOverlap(accumulatedMovieSet, head)) {
        result1 = findPossibleJobsSet(tail, accumulatedMovieSet.concat(head))
    }
    let result2 = findPossibleJobsSet(tail, accumulatedMovieSet)

    return [...result1, ...result2]
}

function findLongestJobsSet(movieSet) {
    const longestJobsSetLength = findPossibleJobsSet(movieSet).reduce((acc, el) => acc < el.length ? el.length : acc, 0)

    return findPossibleJobsSet(movieSet).filter(ms => ms.length === longestJobsSetLength)
}
console.log(findLongestJobsSet(movieSet))

// assert.deepEqual(findLongestJobsSet(movieSet).map(ms=>ms.map(m=>m.name)), [])