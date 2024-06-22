const mysql = require('mysql2/promise');

const pool = mysql.createPool({
    host: process.env.DB_HOST || 'db',
    user: process.env.DB_USER || 'root',
    password: process.env.DB_PASSWORD || 'root',
    database: process.env.DB_DATABASE || 'test',
});

const getUsers = async () => {
    const [rows] = await pool.query('SELECT * FROM user');
    return rows;
};

module.exports = { getUsers };
