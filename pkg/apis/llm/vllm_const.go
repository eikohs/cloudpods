package llm

const (
	LLM_VLLM              = "vllm"
	LLM_VLLM_DEFAULT_PORT = 8000
	LLM_VLLM_EXEC_PATH    = "python3 -m vllm.entrypoints.openai.api_server"

	LLM_VLLM_HF_ENDPOINT = "https://hf-mirror.com"

	// Directory constants
	LLM_VLLM_CACHE_DIR   = "/root/.cache/huggingface"
	LLM_VLLM_BASE_PATH   = "/data/models"
	LLM_VLLM_MODELS_PATH = "/data/models/huggingface"
)
