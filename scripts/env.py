import utils

def copy_env():
    utils.copy_env_file("frontend/.env.example", "frontend/.env")
    utils.copy_env_file("api/.env.example", "api/.env")
    utils.copy_env_file(".env.example", ".env")
    
if __name__ == "__main__":
    copy_env()