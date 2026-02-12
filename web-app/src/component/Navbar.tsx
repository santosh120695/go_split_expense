import { Menu } from "lucide-react";
import { useSidebarStore } from "../store/useStore";

function Navbar() {
  const { toggle } = useSidebarStore();

  return (
    <nav className="bg-(--background) p-4">
      <div className="container mx-auto flex justify-between items-center">
        <button className="md:hidden" onClick={toggle}>
          <Menu />
        </button>
      </div>
    </nav>
  );
}

export default Navbar;
