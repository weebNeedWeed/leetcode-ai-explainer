import React from "react";

const ExplanationDisplay = ({ explanation, isLoading, error }) => {
  if (isLoading) {
    return (
      <div className="bg-white rounded-md p-6 shadow-sm mt-6">
        <div className="animate-pulse">
          <div className="h-6 bg-secondary rounded w-3/4 mb-4"></div>
          <div className="h-4 bg-secondary rounded w-full mb-3"></div>
          <div className="h-4 bg-secondary rounded w-5/6 mb-3"></div>
          <div className="h-4 bg-secondary rounded w-4/6 mb-3"></div>
          <div className="h-4 bg-secondary rounded w-full mb-3"></div>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="bg-white rounded-md p-6 shadow-sm mt-6 border-l-4 border-red-500">
        <h3 className="text-lg font-medium text-red-500 mb-2">Error</h3>
        <p className="text-text-primary">{error}</p>
      </div>
    );
  }

  if (!explanation) {
    return (
      <div className="bg-white rounded-md p-6 shadow-sm mt-6 border-l-4 border-accent">
        <h3 className="text-lg font-medium text-text-primary mb-2">
          No Explanation Yet
        </h3>
        <p className="text-text-secondary">
          Enter a LeetCode question ID above to get an AI-generated explanation.
        </p>
      </div>
    );
  }

  return (
    <div className="bg-white rounded-md p-6 shadow-sm mt-6">
      <h2 className="text-xl font-medium text-text-primary mb-4">
        Solution Explanation
      </h2>
      <div
        className="prose max-w-none text-text-primary"
        dangerouslySetInnerHTML={{ __html: explanation }}
      />
    </div>
  );
};

export default ExplanationDisplay;
