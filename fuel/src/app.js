const express = require('express');

const app = express();
const port = 80;

app.get('/', (req, res) => {
  res.send('I am the goodest fuel service! frick yea!');
});

app.listen(port, () => {
  console.log(`App running on port ${port}`);
});
