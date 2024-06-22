const express = require('express');
const bodyParser = require('body-parser');
const demoRoutes = require('./routes/demoRoutes');

const app = express();

app.use(bodyParser.json());
app.use('/', demoRoutes);

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`);
});
