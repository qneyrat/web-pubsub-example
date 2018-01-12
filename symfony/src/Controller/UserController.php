<?php
declare(strict_types=1);

namespace App\Controller;

use App\Document\Conversation;
use App\Document\Message;
use Doctrine\ODM\MongoDB\DocumentManager;
use Symfony\Bundle\FrameworkBundle\Controller\Controller;
use Symfony\Component\HttpFoundation\JsonResponse;
use Symfony\Component\Routing\Annotation\Route;
use Symfony\Component\Serializer\SerializerInterface;

class UserController extends Controller
{
    /**
     * @var SerializerInterface
     */
    private $serializer;

    /**
     * @var DocumentManager
     */
    private $documentManager;

    /**
     * UserController constructor.
     * @param SerializerInterface $serializer
     * @param DocumentManager $documentManager
     */
    public function __construct(SerializerInterface $serializer, DocumentManager $documentManager)
    {
        $this->serializer = $serializer;
        $this->documentManager = $documentManager;
    }

    /**
     * @Route("/me", methods={"GET"})
     */
    public function meAction()
    {
        $userId = '5a54017b5bb3eb011019f391';

        $conversation = $this->documentManager
            ->getRepository(Conversation::class)
            ->find('5a57c2c85bb3eb002b1f70d1')
        ;

        $message = new Message();
        $message->setFrom($userId);
        $message->setBody('hello world2!');

        $messages = $conversation->getMessages();
        $messages[] = $message;
        $conversation->setMessages($messages);

        $this->documentManager->flush();

        return JsonResponse::fromJsonString(
            $this->serializer->serialize($this->getUser(), 'json', ['groups' => ['user']])
        );
    }
}
