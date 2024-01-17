from typing import List


def validIPAddresses(number_string: str) -> List[str]:
    """
    an IP address has the following structure
    - 4 octets
    - each octet can hold upto 255

    example of the highest IP : 255.255.255.255
    """
    valid_ip_addresses = []
    is_octet_valid = lambda oc: (not int(oc) > 255) and len(oc) == len(str(int(oc)))

    for first_idx in range(1, min(len(number_string), 4)):
        current_ip = ["", "", "", ""]
        current_ip[0] = number_string[:first_idx]
        if not is_octet_valid(current_ip[0]):
            continue
        for second_idx in range(
            first_idx + 1, first_idx + min(len(number_string) - first_idx, 4)
        ):
            current_ip[1] = number_string[first_idx:second_idx]
            if not is_octet_valid(current_ip[1]):
                continue
            for third_idx in range(
                second_idx + 1, second_idx + min(len(number_string) - second_idx, 4)
            ):
                current_ip[2] = number_string[second_idx:third_idx]
                current_ip[3] = number_string[third_idx:]
                if is_octet_valid(current_ip[2]) and is_octet_valid(current_ip[3]):
                    valid_ip_addresses.append(".".join(current_ip))
    return valid_ip_addresses
