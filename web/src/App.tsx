import FileExplorer from "./components/FileExplorer";

function App() {
  return (
    <div className="min-h-screen bg-gray-100 p-4 text-gray-800">
      <h1 className="text-2xl font-bold mb-4">📁 File Share</h1>
      <FileExplorer />
    </div>
  );
}

export default App;
