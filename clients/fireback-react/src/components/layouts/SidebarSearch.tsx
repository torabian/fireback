export default function SidebarSearch() {
  return (
    <form action="/" method="GET" className="search-form">
      <input
        type="search"
        placeholder="Find a setting"
        className="search-field"
      />
      <button type="submit" className="search-button">
        <img src="/adp/icons/mglass.svg" />
      </button>
    </form>
  );
}
