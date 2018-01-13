<?php
declare(strict_types=1);

namespace App\Broker;

use App\Document\Conversation;
use App\Document\Message;
use Swarrot\SwarrotBundle\Broker\Publisher;
use Swarrot\Broker\Message as BrokerMessage;
use Symfony\Component\Serializer\SerializerInterface;

class MessagePublisher
{
    /**
     * @var Publisher
     */
    private $publisher;

    /**
     * @var SerializerInterface
     */
    private $serializer;

    /**
     * MessagePublisher constructor.
     * @param Publisher $publisher
     * @param SerializerInterface $serializer
     */
    public function __construct(Publisher $publisher, SerializerInterface $serializer)
    {
        $this->publisher = $publisher;
        $this->serializer = $serializer;
    }

    public function messageAdded(Conversation $conversation, Message $message)
    {
        foreach($conversation->getUsers() as $user) {
            if($message->getFrom() !== $user) {
                $message->setTo($user);

                $payload = $this->serializer->serialize($message, 'json', ['conversation']);
                $this->publisher->publish(
                    'message',
                    new BrokerMessage($payload),
                    ['routing_key' => sprintf('api.conversation.%s.message.added', $conversation->getId())]
                );
            }
        }
    }
}
