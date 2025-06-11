use strict;
use warnings;
use LWP::UserAgent;
use JSON;

my $url = 'http://localhost:8080/v1/primes?limit=20';
my $ua = LWP::UserAgent->new;
my $res = $ua->get($url);

if ($res->is_success) {
    my $data = decode_json($res->decoded_content);
    print "Primes: ", join(", ", @{$data->{primes}}), "\n";
} else {
    die "HTTP error: " . $res->status_line;
}
