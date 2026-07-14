export default function Header() {
  return (
    <header className="header">
      <div className="header-inner">
        {/* Drop your icon at frontend/public/logo.png to show it here */}
        <div className="header-logo">
          <img src="/logo.png" alt="Logo" onError={(e) => { (e.target as HTMLImageElement).style.display = "none"; }} />
        </div>
        <div>
          <h1>Candidate Screener</h1>
          <p className="header-sub">AI-powered resume screening against your job requirements</p>
        </div>
      </div>
    </header>
  );
}
