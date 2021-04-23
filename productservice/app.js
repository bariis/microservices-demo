const express = require('express');
const app = express();

const port = process.env.PORT || 8080;

app.listen(port, () => {
    console.log(`App is listening on port ${port}`);
});

const port = process.env.PORT || 8080;
