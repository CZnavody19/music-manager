from os import getenv

def get_env_var(name:str) -> str:
    val = getenv(name)
    if not val:
        raise ValueError(f"Environment variable '{name}' is not set.")
    return val
