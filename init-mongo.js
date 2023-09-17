db.createUser({
  user: "blue-server",
  pwd: "123456",
  roles: [{ role: "readWrite", db: "blue-api-local" }],
});

db.createCollection("wallet");

db.wallet.insertOne({
  _id: { $oid: "65045a3eddb9be365f2d8cbb" },
  banknotes: { 20: 10, 50: 10, 100: 10, 500: 10, 1000: 13 },
  total: 19886,
  coins: { 1: 11, 5: 11, 10: 12 },
});

db.createCollection("product");
db.product.insertMany([
  {
    _id: {
      $oid: "6505679b5956f83f47d7dc45",
    },
    name: "water",
    price: 10,
    quantity: 0,
    image: "https://source.unsplash.com/reEySFadyJQ",
    created_at: {
      $date: "2023-09-16T08:30:19.524Z",
    },
    updated_at: {
      $date: "2023-09-16T14:27:23.679Z",
    },
  },
  {
    _id: {
      $oid: "650567b85956f83f47d7dc46",
    },
    created_at: {
      $date: "2023-09-16T08:30:48.190Z",
    },
    updated_at: {
      $date: "2023-09-16T14:33:50.988Z",
    },
    name: "water1",
    price: 10,
    quantity: 94,
    image: "https://source.unsplash.com/dGIEMeN2MV8",
  },
  {
    _id: {
      $oid: "650569165956f83f47d7dc47",
    },
    created_at: {
      $date: "2023-09-16T08:36:38.689Z",
    },
    updated_at: {
      $date: "2023-09-16T08:36:38.689Z",
    },
    name: "water2",
    price: 10,
    quantity: 100,
    image: "https://source.unsplash.com/Aej5gA11eHQ",
  },
  {
    _id: {
      $oid: "650569325956f83f47d7dc48",
    },
    quantity: 100,
    image: "https://source.unsplash.com/h10-NImYZHs",
    created_at: {
      $date: "2023-09-16T08:37:06.481Z",
    },
    updated_at: {
      $date: "2023-09-16T08:37:06.481Z",
    },
    name: "water3",
    price: 10,
  },
  {
    _id: {
      $oid: "650569635956f83f47d7dc49",
    },
    updated_at: {
      $date: "2023-09-16T08:37:55.362Z",
    },
    name: "water4",
    price: 10,
    quantity: 1,
    image: "https://source.unsplash.com/POAQXzBwF7g",
    created_at: {
      $date: "2023-09-16T08:37:55.362Z",
    },
  },
  {
    _id: {
      $oid: "65056acf5956f83f47d7dc4a",
    },
    quantity: 1,
    image: "https://source.unsplash.com/8IQXvtiDmgw",
    created_at: {
      $date: "2023-09-16T08:43:59.695Z",
    },
    updated_at: {
      $date: "2023-09-16T08:43:59.695Z",
    },
    name: "water5",
    price: 10,
  },
  {
    _id: {
      $oid: "65056ade5956f83f47d7dc4b",
    },
    image: "https://source.unsplash.com/Kf6UgCx5mb8",
    created_at: {
      $date: "2023-09-16T08:44:14.943Z",
    },
    updated_at: {
      $date: "2023-09-16T08:44:14.943Z",
    },
    name: "water6",
    price: 10,
    quantity: 1,
  },
  {
    _id: {
      $oid: "65056afc5956f83f47d7dc4c",
    },
    image: "https://source.unsplash.com/T88rQmiCxHs",
    created_at: {
      $date: "2023-09-16T08:44:44.351Z",
    },
    updated_at: {
      $date: "2023-09-16T08:44:44.351Z",
    },
    name: "water7",
    price: 10,
    quantity: 0,
  },
  {
    _id: {
      $oid: "65056b115956f83f47d7dc4d",
    },
    name: "water8",
    price: 10,
    quantity: 0,
    image: "https://source.unsplash.com/40oNjij67qQ",
    created_at: {
      $date: "2023-09-16T08:45:05.919Z",
    },
    updated_at: {
      $date: "2023-09-16T14:28:07.684Z",
    },
  },
]);
