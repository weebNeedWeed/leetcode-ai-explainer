import React from "react";

const Header = () => {
  return (
    <header className="bg-primary py-4 shadow-sm">
      <div className="container mx-auto px-4 max-w-content">
        <div className="flex justify-between items-center">
          <h1 className="text-2xl font-bold text-text-primary">
            LeetCode AI Explainer
          </h1>
          <nav>
            <ul className="flex space-x-6">
              <li>
                <a
                  href="https://github.com/weebNeedWeed/leetcode-ai-explainer"
                  className="text-text-secondary hover:text-accent transition-colors duration-200"
                >
                  GitHub
                </a>
              </li>
            </ul>
          </nav>
        </div>
      </div>
    </header>
  );
};

export default Header;
