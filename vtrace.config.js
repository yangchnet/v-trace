CWD = process.cwd()

module.exports = {
    apps: [
        {
            name: "iam-service",
            script: "make run.iam",
            cwd: CWD,
            watch: [CWD + "/app/iam"],
        },
        {
            name: "ca-service",
            script: "make run.ca",
            cwd: CWD,
            watch: [CWD + "/app/ca"],
        },
        {
            name: "trans-service",
            script: "make run.trans",
            cwd: CWD,
            watch: [CWD + "/app/transaction"],
        },
        {
            name: "goods-service",
            script: "make run.goods",
            cwd: CWD,
            watch: [CWD + "/app/goods"],
        },
        {
            name: "vtrace-service",
            script: "make run.vtrace",
            cwd: CWD,
            watch: [CWD + "/app/vtrace"],
        },
        {
            name: "trans-job",
            script: "make run.trans-job",
            cwd: CWD,
            watch: [CWD + "/app/trans-job"],
        },
        {
            name: "circ-service",
            script: "make run.circ",
            cwd: CWD,
            watch: [CWD + "/app/circ"],
        },
        {
            name: "algo-service",
            script: "make run.algo",
            cwd: CWD,
            watch: [CWD + "/app/algo"],
        },
    ]
}
