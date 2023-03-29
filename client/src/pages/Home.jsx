import React from "react";
import { useEffect } from "react";
import { useState } from "react";
import { Link, useLocation } from "react-router-dom";
import axios from "axios";

const Home = () => {
  const [products, setProducts] = useState([]);
  const [version, setVersion] = useState(0);

  const cat = useLocation().search;

  useEffect(() => {
    const fetchData = async () => {
      try {
        const res = await axios.get(
          `http://localhost:3000/v1/products/all${cat}`,
          { withCredentials: true }
        );
        setProducts(res.data.products);
        console.log(res);
      } catch (err) {
        console.log(err);
      }
    };
    fetchData();
  }, [cat, version]);

  const handleOfferSubmit = (event, productId, offerPrice) => {
    event.preventDefault();
    axios
      .put(
        "http://localhost:3000/v1/products/offer",
        {
          productId: productId,
          offerPrice: parseInt(offerPrice),
        },
        { withCredentials: true }
      )
      .then((response) => {
        console.log(response);
        setVersion((prevVersion) => prevVersion + 1);
      })
      .catch((error) => {
        console.log(error);
      });
  };

  const getText = (html) => {
    const doc = new DOMParser().parseFromString(html, "text/html");
    return doc.body.textContent;
  };

  return (
    <div className="home">
      <div className="products">
        {products.map((product) => (
          <div className="product" key={product.ID}>
            <div className="img">
              <img src={product.ImgUrl} alt="" />
            </div>
            <div className="content">
              <Link className="link" to={`/post/${product.ID}`}>
                <h1>{product.Name}</h1>
              </Link>
              <p>Highest offer: {getText(product.OfferPrice)}</p>
              <form
                onSubmit={(event) => {
                  handleOfferSubmit(
                    event,
                    product.ID,
                    event.target.elements.offerPrice.value
                  );
                }}
              >
                <label htmlFor={`offerPrice_${product.ID}`}>
                  Make an offer:
                </label>
                <input
                  id={`offerPrice_${product.ID}`}
                  type="number"
                  name="offerPrice"
                  min={product.OfferPrice}
                  step="1"
                  required
                />
                <button type="submit">Submit offer</button>
              </form>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Home;
