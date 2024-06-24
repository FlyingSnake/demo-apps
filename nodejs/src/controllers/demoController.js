const { getUsers } = require('../models/user');

const hello = (req, res) => {
    res.send('Hello World (Nodejs / Express)');
};

const users = async (req, res) => {
    try {
        const users = await getUsers();
        res.json(users);
    } catch (error) {
        console.error(error);
        res.status(500).json({ error: 'Database error' });
    }
};

const sleep = (req, res) => {
    const seconds = parseInt(req.params.seconds, 10);
    setTimeout(() => {
        res.send(`sleep ${seconds}s`);
    }, seconds * 1000);
};

const randomStatus = (req, res) => {
    const statuses = [
        { code: 200, message: 'OK' },
        { code: 201, message: 'Created' },
        { code: 202, message: 'Accepted' },
        { code: 204, message: 'No Content' },
        { code: 400, message: 'Bad Request' },
        { code: 401, message: 'Unauthorized' },
        { code: 403, message: 'Forbidden' },
        { code: 404, message: 'Not Found' },
        { code: 500, message: 'Internal Server Error' },
        { code: 501, message: 'Not Implemented' },
        { code: 502, message: 'Bad Gateway' },
        { code: 503, message: 'Service Unavailable' },
    ];
    const randomStatus = statuses[Math.floor(Math.random() * statuses.length)];
    res.status(randomStatus.code).send(randomStatus.message);
};

const exception = (req, res) => {
    const currentTime = new Date().toISOString();
    console.error(`exception called at ${currentTime}`);
    res.status(500).send(`exception called at ${currentTime}`);
};

module.exports = { hello, users, sleep, randomStatus, exception };
