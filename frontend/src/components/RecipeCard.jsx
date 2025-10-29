export default function RecipeCard({ recipe }) {
  return (
    <div className="border rounded-xl p-4 shadow-md bg-white hover:shadow-lg transition">
      <h2 className="text-xl font-bold mb-2">{recipe.name}</h2>
      <p className="text-gray-600 italic">{recipe.country}</p>
      <p className="text-sm mt-2 font-semibold">Type: {recipe.type}</p>
      <p className="text-sm mt-1">{recipe.ingredients}</p>
      <p className="mt-2 text-sm">
        {recipe.is_vegan ? "🌱 Vegan friendly" : "🍖 Non-vegan"}
      </p>
    </div>
  );
}
