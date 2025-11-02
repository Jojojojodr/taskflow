import utils

go_deps = [
    "github.com/air-verse/air@latest",
    "github.com/go-task/task/v3/cmd/task@latest"
]

def install_dependencies():
    print("Installing Go dependencies...\n")
    for dep in go_deps:
        utils.install_go_dependency(dep)  
    print("\nAll Go dependencies installed successfully.")
    
if __name__ == "__main__":
    install_dependencies()