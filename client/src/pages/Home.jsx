import React from "react";
import { useEffect } from "react";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

const Home = () => {
  const [products, setProducts] = useState([]);
  const [version, setVersion] = useState(0);
  const navigate = useNavigate();
  //const cat = useLocation().search;

  useEffect(() => {
    const fetchData = async () => {
      try {
        const res = await axios.get(`/v1/products/all`, {
          withCredentials: true,
        });

        setProducts(res.data.products);
        console.log(res);
      } catch (err) {
        console.log(err);
        if (err.response.status !== 200) {
          navigate("/login");
        }
      }
    };

    const interval = setInterval(() => {
      setVersion((prevVersion) => prevVersion + 1);
    }, 5000);

    fetchData();

    return () => clearInterval(interval);
  }, [version, navigate]);

  const handleOfferSubmit = (event, productId, offerPrice) => {
    event.preventDefault();
    axios
      .put(
        "/v1/products/offer",
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
              <span className="link">
                <h1>{product.Name}</h1>
              </span>
              <div className="highest-offer">
                <span className="highest-offer-header">Highest Offer</span>
                <span className="offer-price">
                  {getText(product.OfferPrice)} ₺
                </span>
                <span className="highest-offer-header">Offered by</span>
                <span className="offerer-name">
                  {getText(product.User.FirstName)}{" "}
                  {getText(product.User.LastName)}
                </span>
              </div>
              <div className="offer-container">
                <form
                  className="offer-form"
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
                    className="offer-input"
                    id={`offerPrice_${product.ID}`}
                    type="number"
                    name="offerPrice"
                    min={product.OfferPrice}
                    step="1"
                    required
                  />
                  <button className="offer-button" type="submit">
                    Submit offer
                  </button>
                </form>
              </div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default Home;
