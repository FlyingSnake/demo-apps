const express = require('express');
const router = express.Router();
const {
    hello,
    users,
    sleep,
    randomStatus,
    exception,
} = require('../controllers/demoController');

router.get('/', hello);
router.get('/users', users);
router.get('/sleep/:seconds', sleep);
router.get('/status/random', randomStatus);
router.get('/exception', exception);

module.exports = router;
