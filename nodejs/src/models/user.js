const mysql = require('mysql2/promise');

const pool = mysql.createPool({
    host: process.env.DB_HOST,
    user: process.env.DB_USERNAME,
    password: process.env.DB_PASSWORD,
    database: process.env.DB_DATABASE
});

const getUsers = async () => {
    const [rows] = await pool.query('SELECT * FROM user');
    return rows;
};

module.exports = { getUsers };
