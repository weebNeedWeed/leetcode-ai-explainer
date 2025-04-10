const fetchExplanation = async (leetCodeId) => {
  if (leetCodeId <= 0) {
    throw new Error("Invalid LeetCode ID");
  }

  const response = await fetch(`/api/problems/${leetCodeId}/explanation`);

  if (!response.ok) {
    throw new Error("Failed to fetch explanation");
  }

  const data = await response.json();

  if (data.code !== 200) {
    throw new Error("Failed to fetch explanation");
  }

  return data.data;
};

export { fetchExplanation };
