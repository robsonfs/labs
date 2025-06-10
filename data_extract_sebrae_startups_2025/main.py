import pdfplumber
import re
import json

FILE_PATH = "local_files/PSS-2025-Top-1.000_06_06-1.pdf"

def data_extract(file_path=None, try_cache=False):
        if try_cache:
            try:
                with open("local_files/startups.json", "r") as cached:
                    data = json.load(cached)
                return data
            except Exception as err:
                print(err)
                print("Attempt to access cached data failed. performing a full parser...")
        data = []
        cnpj_regex = re.compile(r"\d{2}\.\d{3}\.\d{3}/\d{4}-\d{2}")

        with pdfplumber.open(file_path or FILE_PATH) as pdf:
            for page in pdf.pages:
                lines = page.extract_text().split('\n')
                for line in lines:
                    if cnpj_regex.search(line):
                        cnpj = cnpj_regex.search(line).group()
                        parts = line.split(cnpj)
                        if len(parts) == 2:
                            before, after = parts
                            tokens = before.strip().split()
                            if len(tokens) >= 2:
                                uf = tokens[0]
                                startup = " ".join(tokens[1:])
                                categoria = " ".join(after.strip().split()[:-1])
                                estande = after.strip().split()[-1]
                                data.append({
                                    "uf": uf,
                                    "startup": startup,
                                    "cnpj": cnpj,
                                    "categoria": categoria,
                                    "estande": estande
                                })

        return data                        
    

def generate_json(input_data: str | list, output_path=None):
    if type(input_data) not in (str, list):
        raise TypeError(
            f"input_data should be either a str or a dict. Found: {type(input_data)}"
        )
    if type(input_data) == str: 
        input_data = data_extract(input_data)
        
    with open(output_path if output_path else "local_files/startups.json", "w") as fp:
        fp.write(json.dumps(input_data, indent=2, ensure_ascii=False))


def get_stats(data):
    startup_per_uf = {}
    startups_per_category = {}

    for item in data:
        startup_per_uf[item['uf']] = startup_per_uf.get(item['uf'], 0) + 1
        startups_per_category[item['categoria']] = startups_per_category.get(item['categoria'], 0) + 1
    stats = {"per_uf": startup_per_uf, "per_categoria": startups_per_category}

    return stats
