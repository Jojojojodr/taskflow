import subprocess
import utils
import os

def download_tailwind_binary():
    sys, url = utils.get_tailwind_binary_url()
 
    os.makedirs("frontend/bin", exist_ok=True)
    
    if sys == "Windows":
        try:
            subprocess.run(
                ["curl", "-L", url, "-o", "frontend/bin/tailwindcss.exe"],
                check=True
            )
            print("Tailwind CSS binary downloaded successfully.")
        except subprocess.CalledProcessError as e:
            print(f"Failed to download Tailwind CSS binary: {e}")
            raise
    elif sys in ["Linux", "Darwin"]:
        try:
            subprocess.run(
                ["curl", "-L", url, "-o", "frontend/bin/tailwindcss"],
                check=True
            )
            subprocess.run(
                ["chmod", "+x", "frontend/bin/tailwindcss"],
                check=True
            )
            print("Tailwind CSS binary downloaded and made executable successfully.")
        except subprocess.CalledProcessError as e:
            print(f"Failed to download or set permissions for Tailwind CSS binary: {e}")
            raise

if __name__ == "__main__":
    download_tailwind_binary()