import React from "react";

const Footer = () => {
  return (
    <footer className="bg-secondary py-6 mt-12">
      <div className="container mx-auto px-4 max-w-content">
        <div className="flex flex-row md:flex-row justify-center items-center">
          <div className="mb-4 md:mb-0">
            <p className="text-text-secondary text-sm">
              &copy; {new Date().getFullYear()} LeetCode AI Explainer. All
              rights reserved.
            </p>
          </div>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
