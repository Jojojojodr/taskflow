import os
import platform
import tailwind
import go

def main():
    go.install_dependencies()
    
    sys = platform.system()
    if sys == "Windows":
        tailwind_path = "frontend/bin/tailwindcss.exe"
    elif sys in ["Linux", "Darwin"]:
        tailwind_path = "frontend/bin/tailwindcss"
        
    if not os.path.exists(tailwind_path):
        os.makedirs("bin", exist_ok=True)
        tailwind.download_tailwind_binary()
        print(f"Tailwind CSS binary downloaded to {tailwind_path}.")
    else:
        print(f"Tailwind CSS binary already exists at {tailwind_path}.")
        

if __name__ == "__main__":
    main()