import { useEffect, useState } from "react";
import api from "../services/api";
import RecipeCard from "../components/RecipeCard";

export default function Home() {
  const [recipes, setRecipes] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    api
      .get("/recipes")
      .then((res) => setRecipes(res.data))
      .catch((err) => console.error("Error fetching recipes:", err))
      .finally(() => setLoading(false));
  }, []);

  if (loading) {
    return <p className="p-6 text-lg font-semibold">Loading recipes...</p>;
  }

  return (
    <div className="p-6">
      <h1 className="text-3xl font-bold mb-4">🌍 World Recipes</h1>
      {recipes.length === 0 ? (
        <p>No recipes available.</p>
      ) : (
        <div className="grid sm:grid-cols-2 md:grid-cols-3 gap-4">
          {recipes.map((r) => (
            <RecipeCard key={r.id} recipe={r} />
          ))}
        </div>
      )}
    </div>
  );
}
