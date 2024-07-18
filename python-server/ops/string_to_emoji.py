import logging
import os

import torch
from sentence_transformers import SentenceTransformer
from transformers import BartTokenizer, BartForConditionalGeneration


# Model used
# @article{DBLP:journals/corr/abs-2311-01751,
#   author       = {Letian Peng and
#                   Zilong Wang and
#                   Hang Liu and
#                   Zihan Wang and
#                   Jingbo Shang},
#   title        = {EmojiLM: Modeling the New Emoji Language},
#   journal      = {CoRR},
#   volume       = {abs/2311.01751},
#   year         = {2023},
#   url          = {https://doi.org/10.48550/arXiv.2311.01751},
#   doi          = {10.48550/ARXIV.2311.01751},
#   eprinttype    = {arXiv},
#   eprint       = {2311.01751},
#   timestamp    = {Tue, 07 Nov 2023 18:17:14 +0100},
#   biburl       = {https://dblp.org/rec/journals/corr/abs-2311-01751.bib},
#   bibsource    = {dblp computer science bibliography, https://dblp.org}
# }

def _translate_2_emoji(sentence,tokenizer,generator, **argv):
    inputs = tokenizer(sentence, return_tensors="pt").to(device)
    generated_ids = generator.generate(inputs["input_ids"], **argv)
    decoded = tokenizer.decode(generated_ids[0], skip_special_tokens=True).replace(" ", "")
    return decoded

#
logger = logging.getLogger("kotoko.debug")
device = torch.device('cuda' if torch.cuda.is_available() else 'cpu')
# to emoji model
path = "KomeijiForce/bart-large-emojilm"
ebd_model = SentenceTransformer('all-MiniLM-L6-v2')
ebd_model.to(device)
matrix = torch.load(os.path.join("models", "EmojiEmbedding.pt"), map_location=device)
tokenizer = BartTokenizer.from_pretrained(path)
generator = BartForConditionalGeneration.from_pretrained(path)
generator.to(device)

def StringToEmoji(content: str):
    decoded = _translate_2_emoji(content, tokenizer, generator, num_beams=4, do_sample=True,
                                 max_length=100)
    logger.debug(f"Translated emoji: {decoded}")
    return decoded
