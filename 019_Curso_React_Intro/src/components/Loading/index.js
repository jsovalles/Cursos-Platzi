import "./Loading.css"

export function Loading() {
  return (
    <div className="loading-container">
      <div className="spinner"></div>
      <p>Loading your tasks...</p>
    </div>
  );
}
