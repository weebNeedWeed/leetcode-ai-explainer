import { useState } from "react";
import Header from "./components/Header";
import Footer from "./components/Footer";
import LeetCodeForm from "./components/LeetCodeForm";
import ExplanationDisplay from "./components/ExplanationDisplay";
import { fetchExplanation } from "./services/explanationService";

function App() {
  const [explanation, setExplanation] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState("");

  const handleSubmit = async (leetCodeId) => {
    setIsLoading(true);
    setError("");

    try {
      const data = await fetchExplanation(leetCodeId);
      setExplanation(data);
    } catch (err) {
      setError(err.message || "Failed to fetch explanation. Please try again.");
      setExplanation("");
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex flex-col bg-primary">
      <Header />

      <main className="flex-grow container mx-auto px-4 py-8 max-w-content">
        <h1 className="text-3xl font-bold text-text-primary mb-6">
          LeetCode Solution Explainer
        </h1>
        <p className="text-text-secondary mb-8">
          Enter a LeetCode question ID to get an AI-generated explanation of the
          solution.
        </p>

        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
          <div className="lg:col-span-1">
            <LeetCodeForm onSubmit={handleSubmit} isLoading={isLoading} />
          </div>

          <div className="lg:col-span-2">
            <ExplanationDisplay
              explanation={explanation}
              isLoading={isLoading}
              error={error}
            />
          </div>
        </div>
      </main>

      <Footer />
    </div>
  );
}

export default App;
