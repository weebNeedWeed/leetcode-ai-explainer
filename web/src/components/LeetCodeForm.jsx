import React, { useState } from "react";

const LeetCodeForm = ({ onSubmit, isLoading }) => {
  const [leetCodeId, setLeetCodeId] = useState("");
  const [error, setError] = useState("");

  const handleSubmit = (e) => {
    e.preventDefault();

    // Basic validation
    if (!leetCodeId.trim()) {
      setError("Please enter a LeetCode question ID");
      return;
    }

    // Check if input is a valid number
    const id = parseInt(leetCodeId.trim());
    if (isNaN(id) || id <= 0) {
      setError("Please enter a valid LeetCode question ID (positive number)");
      return;
    }

    setError("");
    onSubmit(id);
  };

  return (
    <div className="bg-secondary rounded-md p-6 shadow-sm">
      <h2 className="text-xl font-medium text-text-primary mb-4">
        Enter LeetCode Question ID
      </h2>
      <form onSubmit={handleSubmit}>
        <div className="mb-4">
          <label htmlFor="leetcode-id" className="block text-text-primary mb-2">
            Question ID
          </label>
          <input
            id="leetcode-id"
            type="text"
            value={leetCodeId}
            onChange={(e) => setLeetCodeId(e.target.value)}
            placeholder="e.g., 1, 2, 3"
            className="w-full px-4 py-2 rounded-sm border border-gray-200 focus:outline-none focus:ring-2 focus:ring-accent focus:border-transparent transition-all duration-200"
            disabled={isLoading}
          />
        </div>
        {error && <div className="mb-4 text-red-500 text-sm">{error}</div>}
        <button
          type="submit"
          className={`px-6 py-2 rounded-sm bg-accent text-white font-medium hover:bg-accent/90 transition-colors duration-200 ${
            isLoading ? "opacity-70 cursor-not-allowed" : ""
          }`}
          disabled={isLoading}
        >
          {isLoading ? "Loading..." : "Get Explanation"}
        </button>
      </form>
    </div>
  );
};

export default LeetCodeForm;
