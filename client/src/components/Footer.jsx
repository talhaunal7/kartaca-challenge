import React from "react";
import Logo from "../img/logo.png";

const Footer = () => {
  return (
    <footer>
      <img src={Logo} alt="" />
      <span>
        Made by{" "}
        <a
          target="_blank"
          rel="noopener noreferrer"
          href="https://www.github.com/talhaunal7"
        >
          Talha Ünal
        </a>
      </span>
    </footer>
  );
};

export default Footer;
